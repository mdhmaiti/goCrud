# Basic go and Next js app with the postgres as a db

This is a Crud app

## steps

2. run the docker compose up -d to initialoz the emtpy docker container
3. to be inside the container run `docker exec -it <servicename> psql -U <user> -d <databasename> "docker exec -it goCrud psql -U medhashis -d postgresGoCrud" for my case 
4. The above command gives a interactive terminal with which you can access your container.
5. It will provide an interface like the <databasename>=# now write the command;
6. "\l" list is the previlages and everyting and '\dt' gives the information about the tables
7. Initialize the go; here I am using the "Gorilla" and the "pq" "(this driver is specific for the postges)" driver different database requires different driver.
8. The basic steps for the creating the backend is the same : - " Find a way to connect to the db ", " create the schemas to automatically generate the tables", " make the api end points that is the routers " , " write the logic for the backend"
