## Exercícios:

- **Exercício 1.7**: A chamada de funcao io.Copy(dst, src) lê de src e escreve em dst. Usea no lugar de ioutil.ReadAll
para copiar o corpo da resposta para o.Stdout sem exigir um buffer grande o suficiente para armazenar todo o stream. 
Não esqueca de verificar o resultado de erro de io.Copy.<br><br>
- **Exercício 1.8**: Modifique fetch para o prefixo http:// seja acrecentado a cada URL de argumento, caso esteja faltando.
Você pode usar strings.HasPrefix.<br><br>
- **Exercício 1.9**: Modifique fetch para exibir também o código de status HTTP encontrado em resp.Status.