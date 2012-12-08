//JConf parse a json file with configurations variables and if exists apply
//an inheritance between sections
package jconf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//JSON config file with all configurations files loaded
type JSONConfig struct {
	path         string
	config       map[string]interface{}
	section_json map[string]interface{}
	file_json    map[string]interface{}
}

//Parse new json file and run inheritance logic
func New(file_path string, section string) (*JSONConfig, error) {
	var err error
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}

	v := make(map[string]interface{})

	err = json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	json_section, ok := v[section].(map[string]interface{})
	if !ok {
		return nil, errors.New("Section not found")
	}

	jc := new(JSONConfig)
	jc.path = file_path
	jc.file_json = v
	jc.section_json = json_section

	jc.Load()

	return jc, nil
}

//Load variables on JSON and apply on final config map
func (jc *JSONConfig) Load() {
	if jc.config == nil {
		jc.config = make(map[string]interface{})
	}

	if e, ok := jc.section_json["+"]; ok {
		switch e.(type) {
		case string:
			jc.Override(e.(string))
		case interface{}:
			for _, v := range e.([]interface{}) {
				jc.Override(v.(string))
			}
		}
	}

	for k, v := range jc.section_json {
		jc.config[k] = v
	}
}

//Override section variables on current section
func (jc *JSONConfig) Override(section string) {
	if section, ok := jc.file_json[section].(map[string]interface{}); ok {
		for k, v := range section {
			jc.config[k] = v
		}
	}
}

//Get a variable based on JSONConfig file section
func (jc *JSONConfig) Get(variable string) (interface{}, bool) {
	if value, ok := jc.config[variable]; ok {
		return value, true
	}
	return nil, false
}
