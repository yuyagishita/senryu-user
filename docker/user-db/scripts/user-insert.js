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
    "password": "yagishita"
});

print("_______USER DATA_______");
db.users.find().forEach(get_results);
print("______END USER DATA_____");
