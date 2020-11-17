## Building the Image
```
docker build -t jordan396/kubeinn-frontend .
docker push jordan396/kubeinn-frontend
docker run -it --rm -d -p 8080:80 jordan396/kubeinn-frontend
```