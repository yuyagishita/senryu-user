function get_results(result) {
    print(tojson(result));
}

function insert_customer(object) {
    print(db.customers.insert(object));
}

insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830af"),
    "firstName": "Eve",
    "lastName": "Berger",
    "username": "Eve_Berger",
    "password": "fec51acb3365747fc61247da5e249674cf8463c2",
    "salt": "c748112bc027878aa62812ba1ae00e40ad46d497",
    "addresses": [ObjectId("57a98d98e4b00679b4a830ad")],
    "cards": [ObjectId("57a98d98e4b00679b4a830ae")]
});

print("_______CUSTOMER DATA_______");
db.customers.find().forEach(get_results);
print("______END CUSTOMER DATA_____");
