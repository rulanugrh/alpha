# See documentation
# If you want to deploy by docker-compose you can run `./docker-deploy dockploy`
# Deploy by k8s can run `./docker-deploy kubeploy`

command=$1

if [[ $command == "dockploy" ]]
then
    docker-compose up --env-file .env up -d
    exit 0
else if [[ $command == "kubeploy" ]]
then
    bash kube.sh
    exit
else
    echo "see documentation"
fi