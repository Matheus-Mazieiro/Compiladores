
vars
    x continuous
    y continuous
    z continuous

maximize
    5*x + 2*y + 7*z

subject to

    r1:
        x + y <= 10

    r2:
        x + z <= 8

    r3:
        y + z >= 4

    r4:
        x >= 0
    
    r5:
        y >= 0
      
    r6:
        z >= 0