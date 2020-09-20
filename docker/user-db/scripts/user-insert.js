function get_results(result) {
    print(tojson(result));
}

function insert_user(object) {
    print(db.users.insert(object));
}

insert_user({
    "_id": ObjectId("5f4e6054ee11cb011220ca4a"),
    "username": "test",
    "password": "c80993c9a9ca4ca914576a54f6876d89fae2d17b",
    "email": "test@test.com",
    "salt": "e3ae274204fc6ef7f0b60b5feca1b049b2d824bb"
});

print("_______USER DATA_______");
db.users.find().forEach(get_results);
print("______END USER DATA_____");
