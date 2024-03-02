# Basic go and Next js app with the postgres as a db

This is a Crud app

## steps

2. run the docker compose up -d to initialoz the emtpy docker container
3. to be inside the container run `docker exec -it <servicename> psql -U <user> -d <databasename> `
4. The above command gives a interactive terminal with which you can access your container.
5. It will provide an interface like the <databasename>=# now write the command;
6. "\l" list is the previlages and everyting and '\dt' gives the information about the tables
