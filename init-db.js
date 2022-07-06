db.createUser(
  {
      user: "ery",
      pwd: "erySRG",
      roles: [
          {
              role: "readWrite",
              db: "gocommerce-db"
          }
      ]
  }
);