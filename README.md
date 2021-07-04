# Nomad & Lightrun

This project is intended for demonstrating how Lightrun can easily be
integrated with Nomad

- Website: [https://www.nomadproject.io](https://www.nomadproject.io)
- Mailing list: [Google Groups](http://groups.google.com/group/nomad-tool)

## Building The Plugin

If needed, one can build the plugin using the instructions bellow

### Requirements For Building The Plugin

- [Nomad](https://www.nomadproject.io/downloads.html) v0.9+
- [Go](https://golang.org/doc/install) v1.11 or later (to build the plugin)

### Building the Driver

Clone the repository somewhere in your computer. This project uses
[Go modules](https://blog.golang.org/using-go-modules)

```sh
$ git clone https://github.com/lightrun-platform/lightrun-n-nomad.git
```

Build the plugin using make:

```sh
$ make build
```

This will run the Go compiler and will eventually create 'lightrun-java-driver' 
binary in your cloned directory

## Using Lightrun With Nomad

When coming to use Lightrun with Nomad, there are 2 possibilities:
1. Use Lightrun as a driver - Lightrun presents a unique driver (same as the 'java' driver )
that does all the magic of integrating with Lightrun for you.
   
2. Use Lightrun as standalone - You can choose to download Lightrun's agent manually
and run your task with it.

### Driver Integration

1. Copy Lightrun's plugin (either built by you or downloaded) to the directory Nomad's
   agent is running from.
2. When running Nomad's agent, you need to specify the plugin's directory:
```sh
sudo nomad agent -dev -bind 0.0.0.0 -log-level DEBUG -plugin-dir=<path_to_plugin_directory>
```
3. See our example of Nomad's configuration found at `example/example.driver.nomad` or simply change
the driver to `lightrun-java` on your job and pass the right configuration to your java application. 
   You can also use the plugin's configuration found at `example/agent.hcl` but then you'll have to 
   run the agent with that configuration:
```sh
sudo nomad agent -dev -bind 0.0.0.0 -log-level DEBUG -config=./agent.hcl -plugin-dir=<path_to_plugin_directory>
```
4. Running the job can be done by `nomad job run ./example.driver.nomad`
   
That's it! Lightrun's plugin will use the arguments given on the config and will automatically
download Lightrun's agent and use it.

### Standalone Integration

1. Use the example found at `example/example.standalone.nomad`. The example downloads 
Lightrun's agent and then uses it with your Jar.
2. `example/example.standalone.nomad.vars` holds the configuration needed by the job.
In this configuration you can find:
   1. `lightrun-server` - That contains the server's address including the company's name
   2. `lightrun-secret` - A secret for the agent to identify itself against the server
   3. `lightrun-certificate` - A thumbprint of the server's certificate
3. Running the job with the configuration is done by `nomad job run -var-file=./example.standalone.nomad.vars ./example.standalone.nomad`