
sets
    V

vars
    x[i,j] binary

subject to

subtour(S subsetof V):

    2 <= |S| <= |V|-1

    sum(i in S, j in V-S)
        x[i,j]
    >= 1

end