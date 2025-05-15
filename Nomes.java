import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;

public class Nomes {
    public static void main(String[] args) {
        // Criando a lista de nomes
        ArrayList<String> nomes = new ArrayList<>();
        nomes.add("Giovanna");
        nomes.add("Mateus");
        nomes.add("Hiro");
        nomes.add("Gilberto");

        // Adicionando mais um nome
        nomes.add("Samira");

        // Escrevendo os nomes no arquivo (modo append)
        try {
            FileWriter escritor = new FileWriter("names.txt", true); // true = modo append
            for (String nome : nomes) {
                escritor.write(nome + "\n");
            }
            escritor.close();
            System.out.println("Nomes adicionados com sucesso!");
        } catch (IOException e) {
            System.out.println("Ocorreu um erro ao escrever no arquivo.");
            e.printStackTrace();
        }
    }
}
