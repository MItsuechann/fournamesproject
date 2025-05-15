
nomes = ["Giovanna", "Mateus", "Hiro", "Gilberto"]


nomes.append("Samira")


with open("names.txt", "a", encoding="utf-8") as arquivo:
    for nome in nomes:
        arquivo.write(nome + "\n")