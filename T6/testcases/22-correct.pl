sets
    I
    J

params
    c[I]
    d[I,J]

vars
    x[I] binary
    y[I,J] continuous

maximize
    sum(i in I)
        x[i]

subject to

cap(i in I):
    x[i] <= c[i]

flow(i in I, j in J):
    y[i,j] >= d[i,j]