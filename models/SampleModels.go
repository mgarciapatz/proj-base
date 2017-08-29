package models

/*
    NOTE
        Structures that are declared in this area is for
        database model mapping purposes only.
*/

type SampleData struct {
    Autokey         int `orm:"pk" json:"-"`
    SampleFirstName string `json:"sample_first"`
    SampleLastName  string `json:"sample_last"`
}
