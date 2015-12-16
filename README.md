shudi
=============

Sometimes, we want to block programs from running - but we don't want to do something drastic like "rename the binary".

We could achieve this by running a small program that:

1. Runs with a schedule in mind - and a random splay amount.
2. Checks to see if program execution should be stopped. (By looking inside of Consul's KV store - for example.)
3. Executing if a block isn't present.
4. Then looping for the appointed time.

The program I'd like to try this with is Chef.

1. We can't stop a Chef run *during* the run without throwing an error. This is not an error.
2. We often don't want Chef to run for all sorts of reasons.
3. The methods we have to stop Chef from running are not great.

I'd like to make the `block` interface simple:

`./shudi block chef-client`

Also:

`./shudi unblock chef-client`

TODO:

1. Add the ability to check Consul's KV store for a key.
2. Add the ability to block this node from running a specific program through `shudi`.
3. Add metrics around operation.
