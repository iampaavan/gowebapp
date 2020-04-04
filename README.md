# gowebapp

API 1 --> http://localhost:8086
API 2 --> http://localhost:8086/time?tz1=EST

Can also try http://localhost:8086/time?tz1=Local&tz2=EST to display all timezones in requests. WIP !!

git clone https://github.com/udhos/update-golang
cd update-golang
sudo ./update-golang.sh

Updated:

1) Bring up your kubernetes cluster
2) Make sure the latest go-webapp image is pushed to DockerHub via Jenkins CI\CD pipeline
3) Go to k8s folder and install the necessary
4) Verify the status of your kubernetes resources in the Dashboard
5) Test the following APIs:
   a) API 1 --> http://localhost:8086 --> Gives to the current local time in your current location
   b) API 2 -- > http://localhost:8086/time?tz1=Local&tz2=EST --> Gives you both the timezones mentioned in the URL
   c) API 3 --> http://locahost:8086/health --> TO verify the application up and running
   d) API 4 --> http://locahost:8086/readiness --> To verify the application is ready to accept requests