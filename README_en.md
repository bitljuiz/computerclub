## computerclub

computerclub - a prototype of a system that monitors the work of a computer club and processes events
and calculates revenue for the day and time of occupancy of each table.

For each package, its detailed description is available in `README_en.md` files, as well as in the form of comments in the code

## Launch

Devices running under Windows (via WSL 2) and Linux are suitable for launching.
You also need to have the following installed on your device:
- Version `go` is not earlier than 1.19 (The project is written on version 1.22.2)
- docker

After cloning the repository, go to the root of the project (the `/computerclub` folder) and run:

```
docker build -t computerclub .
docker run --rm computerclub <filename>
```

This command starts docker, builds the project and outputs the result of the work in accordance with the input data.
In place of `<filename>`, respectively, indicate the path to the file in which the input data is located.

If you have problems running the project, try restarting docker through the application or with the command:

```
sudo systemctl restart docker
```

If this doesn't help, make sure you are logged into docker using the command:

```
docker login
```

To run a test from a condition, instead of `<filename>` enter `test_file.txt` (the file is located in the project root)