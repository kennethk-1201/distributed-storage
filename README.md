# Distributed File Storage (DFS)

The distributed file storage **in the context of this repository** is a shared network of multiple
users replicating their files over other devices in the network. Hence, this distributed file
storage aims to achieve **availability** of files stored locally.

## Writes
When writing a new file to the DFS, the host will replicate this file to all its peers for durability.

## Reads
When reading a new file from the DFS, the host will check its local filesystem first. If the file
does not exist, it will retrieve file A from some other peer in the network.

## Encryption
Each host has its own ID. Each file will be related to an ID, which means that this file belongs to
the corresponding host with the same ID.

If the ID of the file belongs to the file's host, the data will be unencrypted as there is no need
to protect the files from its owner. Otherwise, it means the file is replicated and will thus
have to be encrypted to protect it from the host.

# Source
This project is meant for my personal learning in Golang and distributed file systems. I referenced
this video for this project: https://www.youtube.com/watch?v=bymQakvTY40
