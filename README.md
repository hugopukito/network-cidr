# IP Overlap

## Requirements
Production-ready Go CLI that prints to STDOUT the relation between two CIDRs. The relations can be:

- **subset**: if the network of the second address is included in the first one
- **superset**: if the network of the second address includes the first one
- **different**: if the two networks are not overlapping
- **same**: if both address are in the same network

The program is only intended to work with IPv4 addresses.

## Examples

## Usage examples

**$> ./overlap 10.0.0.0/20 10.0.2.0/24**
subset

**$> ./overlap 10.0.2.0/24 10.0.0.0/20**
superset

**$> ./overlap 10.0.2.0/24 10.0.3.0/24**
different

**$> ./overlap 10.0.2.0/24 10.0.2.10/24**
same