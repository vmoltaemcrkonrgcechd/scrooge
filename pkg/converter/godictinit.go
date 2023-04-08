package converter

func (c *Converter) goDictInit() {
	c.goDict = make(map[string]string)

	c.goDict["smallint"] = "int"
	c.goDict["integer"] = "int"
	c.goDict["bigint"] = "int"
	c.goDict["smallserial"] = "int"
	c.goDict["serial"] = "int"
	c.goDict["bigserial"] = "int"

	c.goDict["decimal"] = "float64"
	c.goDict["numeric"] = "float64"
	c.goDict["real"] = "float64"
	c.goDict["double precision"] = "float64"

	c.goDict["character varying"] = "string"
	c.goDict["character"] = "string"
	c.goDict["text"] = "string"
	c.goDict["uuid"] = "string"

	c.goDict["boolean"] = "bool"
}
