# sum: Do the math with multiple workers

Let's say we want to calculate 1+2+...+10000(just pretend you have no idea how to work it out), and
we want to do it channel way.

we split the work loads among workers, and let them calculate part of the whole job, and sum the result
to get the final number.

NOTE: CPU-heavy tasks does not benefit much from channels, because all workers still need the same amount of CPU time.
And no one is doing any IO operatios while the other is busy. 
