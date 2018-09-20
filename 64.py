sw = "switch n {\n"
for i in range(1,65):
    sw = sw+f"""case n<{1<<i}&&n>={1<<(i-1)}:
        return {i-1}
    """
sw += "}"
print(sw)

