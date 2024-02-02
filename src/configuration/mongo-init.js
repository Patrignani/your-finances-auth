let today = new Date();
const dbName = 'your-finances-auth'
const db = db.getSiblingDB(dbName);
const collections = db.getCollectionNames();

function upsertDocument(db, collection, filter, document, dbName) {
    let update = { $set: document};
    let options = { upsert: true};
    db.getSiblingDB(dbName)[collection].updateOne(filter, update, options);
  }

function CreateCollection(collections, db, collectionName){   
    if (!collections.includes(collectionName)) {
      db.createCollection(collectionName);
    }
  }

db.createUser(
    {
      user: 'clienteOcultoAuthAdmin',
      pwd: 'f0cd47b4b7364a7e9b87e1a377b7dddf',
      roles: [{ role: 'readWrite', db: dbName }],
    },
  );

  CreateCollection(collections, db, 'clients')
  CreateCollection(collections, db, 'roles')
  CreateCollection(collections, db, 'users')
  CreateCollection(collections, db, 'permissions')
  CreateCollection(collections, db, 'refresh_tokens')

//clients index
db.clients.createIndex({client_id: 1}, {unique: true});
db.clients.createIndex({client_id: 1, client_secret: 1});

//roles index 
db.roles.createIndex({code: 1}, {unique: true});
db.roles.createIndex({name: 1}, {unique: true});

//users index
db.users.createIndex({ username: 1 }, { unique: true });
db.users.createIndex({ username: 1, active: 1});

//pemissions index
db.permissions.createIndex({code: 1}, {unique: true});
db.permissions.createIndex({name: 1});

//refresh token index
db['refresh_tokens'].createIndex({_id: 1, active:1, expiration_date:1});

//clients data initial
let client = {
  client_id: '2f6f931db8e84b179de0f34f278c977f',
  client_secret: '0ddfc869916b4130ad804340ba2f7cdb',
  name: 'front-end-vue',
  description: 'fronte end vue.js',
  create_at: today,
  update_at: today,
  create_by: 'admin-docker',
  client_create_by: 'admin-docker',
  active: true,
  roles: [],
  permissions: []
};

let clientApi = {
  client_id: '10562f253bdb465397e613e28e356878',
  client_secret: 'e784ac6efdb84bcc82d670aebc729800',
  name: 'your-finances-api',
  description: 'Client da API',
  create_at: today,
  update_at: today,
  create_by: 'admin-docker',
  client_create_by: 'admin-docker',
  active: true,
  roles: [],
  permissions: []
};

upsertDocument(db, "clients", {client_id: client.client_id}, client, dbName);
upsertDocument(db, "clients", {client_id: clientApi.client_id}, clientApi, dbName);

//roles data initial
let role = {
  code: 1,
  name: 'admin',
  description: 'role para administrador'
};
upsertDocument(db, "roles", {code: role.code}, role, dbName);

let user = {
  username: "admin",
  password: "$2a$10$AJn9kH6hRJ8wFUYDgBDnieQItckbzJybXXN8NknX/kYfc7As1VQyO",
  seed: "213ca3d614b04698a94068afa45672cc",
  roles: ["1"],
  permissions: [],
  create_at: today,
  create_by: "admin",
  client_create_by: "admin",
  active: true
}
upsertDocument(db, "users", {username: user.username}, user, dbName);

let user2 = {
    username: "Patrignani",
    password: "$2a$10$MPXj2K4cE6RgtUs8Ik1ib.XuV0ZISvv7bv8mf9Rzyte2VjRWl.CWe",
    seed: "ca2f66de4efa454287d4ca1cf989dd31",
    roles: [],
    permissions: [],
    create_at: today,
    create_by: "admin",
    client_create_by: "admin",
    active: true,
    two_factory_code: "FQ5J6TV4FV4BOEPU7TNRXVAATFTRZQNJ"
  }
  upsertDocument(db, "users", {username: user2.username}, user2, dbName);

  let user3 = {
    username: "master",
    password: "$2a$10$Ia.Q0fuOD0GUA8o8Ml8/seMfCJmctG5Xj1qYqBA4EX5LvRwnEolFu",
    seed: "bf3ef5ce40f34a4381b7d27f894191d6",
    roles: [],
    permissions: [],
    create_at: today,
    create_by: "admin",
    client_create_by: "admin",
    active: true
  }
  upsertDocument(db, "users", {username: user3.username}, user3, dbName);



