# SimpleQ

This is a very simple Queue that reads a text file with payloads and executes the script using the payloads asynchronously. You can specifiy how many tasks can be executed in parallel.

The tool has three parameters: 
* script (string): the script to be executed with the payload from the queue (make sure it's executable)
* queue (string): the queue file to be watched
* maxJobs (int): the max # of processes that can be spun up in parallel

## How to use
```
./simpleQ -script foobar.bash -queue queue/sample -maxJobs 4
```

To be used in [BrainImAccs/brainSTEM](https://github.com/BrainImAccs/BrainSTEM)