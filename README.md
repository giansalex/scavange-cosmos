# Scavenge - cosmos

## Run Node
```bash
make
./init.sh
scavengeD start
```

Other terminal (client commands)
```bash
# show me account balance
scavengeCLI q account $(scavengeCLI keys show me -a)
# show you account balance
scavengeCLI q account $(scavengeCLI keys show you -a)

# start
scavengeCLI tx scavenge createScavenge 10foo "giansalex" "who am i?" --from me
# commit solution
scavengeCLI tx scavenge commitSolution "giansalex" --from you
# verify commited solution
scavengeCLI q scavenge commited "giansalex" $(scavengeCLI keys show you -a)
# reveal solution
scavengeCLI tx scavenge revealSolution "giansalex" --from you

# show me account balance
scavengeCLI q account $(scavengeCLI keys show me -a)
# show you account balance
scavengeCLI q account $(scavengeCLI keys show you -a)

```


Run REST client
```
scavengeCLI rest-server --chain-id scavenge --trust-node
```

## Run second node

Get node id (First machine)
```bash
scavengeD tendermint show-node-id
```

Init node 2 
```
scavengeD init mynode2 --chain-id scavenge
```

Configure peers
```bash
# Replace NODE_1_ID
vim /.nsd/config/config.toml
persistent_peers = "NODE_1_ID@first_node_ip:26656"
```


Start node 2
```
scavengeD start
```