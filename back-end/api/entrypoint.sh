# wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

echo "------------------"
env | grep DATABASE
echo "------------------"

#mysql -h mysql-djb -u root -p"lmi123" < ./back-end/init.sql

cd ./back-end/api/cmd/main

#go run main.go
sleep 10000
