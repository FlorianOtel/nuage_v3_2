package nuage_v3_2

import (
	"encoding/json"
	"fmt"

	nuage "github.com/FlorianOtel/nuage"

	log "github.com/Sirupsen/logrus"
)

func (orglist *EnterpriseSlice) EnterprisesList(c *nuage.Connection) error {

	log.Debug("Entering: EnterpriseList")

	// XXX - Alternative
	// var orgs []Enterprise

	reply, err := nuage.ListEnterprises(c)
	if err != nil {
		log.Debugf("EnterpriseList: Unable to obtain Enterprise list: %s ", err)
		return err
	}

	// XXX - Alternative
	// err = json.Unmarshal(reply, &orgs)
	err = json.Unmarshal(reply, orglist)

	if err != nil {
		log.Debugf("EnterpriseList: Unable to decode JSON payload: %s ", err)
		return err
	}

	//// var orgs []Enterprise
	////
	//// orgs = *orglist
	////
	//// for i, v := range orgs {
	//// 	jsonorg, _ := json.MarshalIndent(v, "", "\t")
	//// 	fmt.Printf("\n\n ===> Org nr [%d]: [%s] <=== \n%#s\n", i, orgs[i].Name, string(jsonorg))
	//// }

	// XXX - Alternative: This effeticvely converts "EnterpriseSlice" to "[]Enterprise"
	// *orglist = orgs

	// log.Fatal("\n\n KABOOM ?? Yes Rico, KABOOM..\n\n")

	return nil
}

// Enterprise Delete
func (org *Enterprise) EnterpriseDelete(c *nuage.Connection) error {

	// XXX -- this will mutate the receiver

	err := org.EnterpriseGet(c)

	if err != nil {
		log.Debugf("EnterpriseDelete: Unable to delete Enterprise with name %s . Error: %s ", org.Name, err)
		return err
	}

	_, err = nuage.DeleteEnterprise(c, org.ID)

	if err != nil {
		log.Debugf("EnterpriseDelete: Unable to delete Enterprise with name %s . Error: %s ", org.Name, err)
		return err
	}

	log.Debugf("EnterpriseDelete: Deleteing Enterprise [%s] with ID: [%s] ", org.Name, org.ID)

	return nil

}

// Assumes that the method receiver was allocated using "new(Enterprise)", initialized accordingly (name + description).  Relies on lower level / API version indep ListEnterprises
func (org *Enterprise) EnterpriseCreate(c *nuage.Connection) error {

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var myorg [1]Enterprise

	myorg[0].Name = org.Name

	if org.Description == "" {
		// Default Enterpise Description unless one is specified
		myorg[0].Description = "Created by Golang API driver"
	} else {
		myorg[0].Description = org.Description
	}

	jsonorg, _ := json.MarshalIndent(myorg[0], "", "\t")

	// Quick and dirty alternative: Just build a JSON object with "name" and "description" fields
	// jsonorg := "      {\"name\":\"" + name + "\",\"description\":\"Created by Golang API client\"}      "

	reply, err := nuage.CreateEnterprise(c, jsonorg)

	if err != nil {
		log.Debugf("EnterpriseCreate: Unable to create Enterprise with name: %s . Error: %s ", org.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &myorg)

	if err != nil {
		log.Debugf("EnterpriseCreate: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX -- Mutate the method receiver
	*org = myorg[0]
	return nil

}

// Assumes that the method receiver was allocated using "new(Enterprise)" with name (org.Name) or ID (org.ID) initialized.  Relies on lower level / API version indep ListEnterprises
func (org *Enterprise) EnterpriseGet(c *nuage.Connection) error {

	reply, err := nuage.ListEnterprises(c)

	if err != nil {
		log.Debugf("EnterpriseGet: Unable to obtain Enterprise list: %s ", err)
		return err
	}

	// XXX -- Need an []Enterprise since the answer is an array of JSON objects (with a single member)
	var orgs []Enterprise
	err = json.Unmarshal(reply, &orgs)

	if err != nil {
		log.Debugf("EnterpriseGet: Unable to decode JSON payload: %s ", err)
		return err
	}

	for _, v := range orgs {
		if org.Name == v.Name || org.ID == v.ID {
			// XXX -- Mutate the method receiver
			*org = v
			return nil
		}
	}

	err = fmt.Errorf("EnterpriseGet: Unable to find Enterprise with name: %s", org.Name)
	return err

}
