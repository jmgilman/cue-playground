package main

#Schema: {
    foo: #Foos
}

#Foos: [string]: #Bar

#Bar: {
    value: string
}