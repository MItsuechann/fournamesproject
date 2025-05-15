const fs = require('fs');

let nomes = ['Giovanna', 'Mateus', 'Miro', 'Gilberto'];

// Adicionando mais um nome
nomes.push('Samira');

// Conteúdo formatado
const conteudo = nomes.join('\n') + '\n';

// Escrevendo no arquivo 'names.txt' em modo append
fs.appendFile('names.txt', conteudo, (err) => {
    if (err) {
        console.error('Erro ao escrever no arquivo:', err);
    } else {
        console.log('Nomes adicionados com sucesso!');
    }
});
