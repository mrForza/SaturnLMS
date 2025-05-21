db = db.getSiblingDB('study');

db.createUser({
    user: 'appuser',
    pwd: 'password',
    roles: [{
        role: 'readWrite',
        db: 'study'
    }]
});

function isUUID(str) {
    return str.match(/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i);
}

function isArrayOfType(arr, typeCheck) {
    if (!Array.isArray(arr)) return false;
    for (var i = 0; i < arr.length; i++) {
        if (!typeCheck(arr[i])) return false;
    }
    return true;
}

function isArrayOfStrings(arr) {
    return isArrayOfType(arr, function(x) { return typeof x === 'string'; });
}

function isArrayOfUUIDs(arr) {
    return isArrayOfType(arr, function(x) { return typeof x === 'string' && isUUID(x); });
}

function isArrayOfFiles(arr) {
    if (!Array.isArray(arr)) return false;
    for (var i = 0; i < arr.length; i++) {
        var f = arr[i];
        if (typeof f.Id !== 'string' || !isUUID(f.Id) || typeof f.Name !== 'string') {
            return false;
        }
    }
    return true;
}

db.createCollection("courses", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["Id", "Name", "Description", "Formula", "Languages", "Teachers", "Students"],
      properties: {
        Id: {
          bsonType: "string",
          description: "должен быть UUID"
        },
        Name: {
          bsonType: "string",
          description: "имя курса"
        },
        Description: {
          bsonType: "string"
        },
        Formula: {
          bsonType: "string"
        },
        Languages: {
          bsonType: "array",
          items: {
            bsonType: "string"
          }
        },
        Teachers: {
          bsonType: "array",
          items: {
            bsonType: "string"
          }
        },
        Students: {
          bsonType: "array",
          items: {
            bsonType: "string"
          }
        }
      }
    }
  }
});

db.createCollection("lessons", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["Id", "Name", "Description", "Type", "Pictures", "Documents"],
      properties: {
        Id: {
          bsonType: "string",
          description: "должен быть UUID"
        },
        Name: {
          bsonType: "string"
        },
        Description: {
          bsonType: "string"
        },
        Type: {
          bsonType: "bool"
        },
        Pictures: {
          bsonType: "array",
          items: {
            bsonType: "object",
            required: ["Id", "Name"],
            properties: {
              Id: { bsonType: "string" },
              Name: { bsonType: "string" }
            }
          }
        },
        Documents: {
          bsonType: "array",
          items: {
            bsonType: "object",
            required: ["Id", "Name"],
            properties: {
              Id: { bsonType: "string" },
              Name: { bsonType: "string" }
            }
          }
        }
      }
    }
  }
});

db.createCollection("homeworks", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["Id", "Name", "Description", "Files"],
      properties: {
        Id: {
          bsonType: "string",
          description: "должен быть UUID"
        },
        Name: {
          bsonType: "string"
        },
        Description: {
          bsonType: "string"
        },
        Files: {
          bsonType: "array",
          items: {
            bsonType: "object",
            required: ["Id", "Name"],
            properties: {
              Id: { bsonType: "string" },
              Name: { bsonType: "string" }
            }
          }
        }
      }
    }
  }
});