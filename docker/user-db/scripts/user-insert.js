function get_results(result) {
    print(tojson(result));
}

function insert_user(object) {
    print(db.users.insert(object));
}

insert_user({
    "_id": ObjectId("57a98d98e4b00679b4a830af"),
    "firstName": "yu",
    "lastName": "yagishita",
    "username": "yagiyu",
    "password": "fec51acb3365747fc61247da5e249674cf8463c2",
    "salt": "c748112bc027878aa62812ba1ae00e40ad46d497"
});

print("_______USER DATA_______");
db.users.find().forEach(get_results);
print("______END USER DATA_____");
