#! /bin/sh

case ${1} in
    "build")
        docker image remove cwb_crawl
        docker image remove web_forecast
 
        docker build -f docker/CrawlerDockerfile -t cwb_crawl .
        docker build -f docker/ForecastDockerfile -t web_forecast .

        # delete <none> images
        docker rmi -f $(docker images --filter "dangling=true" -q)
        docker images
    ;;
    "start")
        docker-compose -f docker/docker-compose.yml up --build
    ;;
    "stop")
        docker-compose -f docker/docker-compose.yml down
    ;;
    "logs")
        docker-compose -f docker/docker-compose.yml logs -f
    ;;
    *)
        docker-compose -f docker/docker-compose.yml ps
    ;;
esac