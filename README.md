This is a simple app that prints the hostname. Useful for testing container load balancing etc.

Container images you can use:

**Node.js**
```
docker run -d -p 8000:8000 oguzpastirmaci/hostname
```

**Go**
```
docker run -d -p 8000:8000 oguzpastirmaci/hostname:go
```

Then browse the URL:8000 of your server/container/pod.
