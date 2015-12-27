shudi
=============

## shudi - pronounced "Should I?"

Sometimes, we want to block programs from running - but we don't want to do something drastic like "rename the binary".

We could achieve this by running a small program that:

1. Runs with a schedule in mind - and a random splay amount.
2. Checks to see if program execution should be stopped. (By looking inside of Consul's KV store - for example.)
3. Executing if a block isn't present.
4. Then looping for the appointed time.

The program I'd like to try this with is Chef client:

`./shudi run -e chef-client -d 1800 -s 900`

1. We can't stop a Chef run *during* the run without throwing an error. This would not throw an error - it would log a metric if it was blocked.
2. We often don't want Chef to run for all sorts of reasons.
3. The methods we have to stop Chef from running are currently not that great.

I'd like to make the `block` interface simple:

`./shudi block -e chef-client`

Also:

`./shudi unblock -e chef-client`
