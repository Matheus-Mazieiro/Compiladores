
sets
    I
    J

vars
    x[i,j] binary

subject to

    forall(i in I)

        assign:

            sum(j in J)
                x[i,j]
            <= 1
