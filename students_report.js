let students = 
[
    {
    "Name":"Lee",
    "RollNo":6,
    "Age":25,
    "Placed":true,
    "Company":{
        "name":"Dell",
        "CTC": 7.5
    },
    "Skills":["NodeJs", "React","MernJS"]

    },
    {
    "Name":"Dee",
    "RollNo":66,
    "Age":25,
    "Placed":false,
    "Company":{
        "name": null,
        "CTC": null
    },
    "Skills":["HTML", "CSS","MernJS"]

    }
]
let stringifiedstudents = JSON.stringify(students)
let parsedstudents = JSON.parse(stringifiedstudents)
parsedstudents.forEach(students => {
    console.log('=========================')
    console.log(students)
})