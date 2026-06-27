
sets
    I
    J

vars
    x[i,j] binary

maximize
    sum(i in I, j in J)
        profit[i,j] * x[i,j]

subject to

    task:
        sum(i in I)
            x[i,j]
        <= 1

    worker:
        sum(j in J)
            x[i,j]
        <= 2

end