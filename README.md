# Snowflake
## Description
There is a 64-bit integer coposed of 3 parts
1. Unused(1 bit)
2. Timestamp (41 bits)
3. Machine ID (10 bits)
4. Sequence Number (12 bits)

## How to get Machine ID
1. Maintain a table with column MachineId, UpdateTime
2. Get minimum number from the table when service is start
3. Regularly update the UpdateTime on this table
4. Regularly delete expired MachineId