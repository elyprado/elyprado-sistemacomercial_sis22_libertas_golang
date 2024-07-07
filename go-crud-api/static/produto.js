idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;
    const txtdescricao = document.getElementById("txtdescricao");
    const txtprecocusto = document.getElementById("txtprecocusto");
    const txtprecovenda = document.getElementById("txtprecovenda");
    const txtsaldoestoque = document.getElementById("txtsaldoestoque");
    const txtcodbarras = document.getElementById("txtcodbarras");
    const txtidmarca = document.getElementById("txtidmarca");
     

    //limpa os campo
    txtdescricao.value = "";
    txtprecocusto.value = "";
    txtprecovenda.value = "";
    txtsaldoestoque.value = "";
    txtcodbarras.value = "";
    txtidmarca.value = "";

    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;
    //carregar os dados do id passado por parametro
    fetch("http://127.0.0.1:8080/produto/" + id)
    .then(resp => resp.json())
    .then(dados => {
        //preenche os inputs
        const txtdescricao = document.getElementById("txtdescricao");
        const txtprecocusto = document.getElementById("txtprecocusto");
        const txtprecovenda = document.getElementById("txtprecovenda");
        const txtsaldoestoque = document.getElementById("txtsaldoestoque");
        const txtcodbarras = document.getElementById("txtcodbarras");
        const txtidmarca = document.getElementById("txtidmarca");


        txtdescricao.value = dados.descricao;
        txtprecocusto.value = dados.precocusto;
        txtprecovenda.value = dados.precovenda;
        txtsaldoestoque.value = dados.saldoestoque;
        txtcodbarras.value = dados.codbarras;
        txtidmarca.value = dados.idmarca;
        

        //mostra a dialog para alterar
        modal.show();
    });
}
function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    fetch("http://127.0.0.1:8080/produto?pesquisa=" + txtpesquisa.value)
    .then(resp => resp.json())
    .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    //limpa a lista
    lista.innerHTML = "";
    //percorre os dados
    for (var i in dados) {
        let id = dados[i].idproduto;
        lista.innerHTML += "<tr>"
            + "<td>" + id + "</td>"
            + "<td>" + dados[i].descricao + "</td>"
            + "<td>" + dados[i].precocusto + "</td>"
            + "<td>" + dados[i].precovenda + "</td>"
            + "<td>" + dados[i].saldoestoque + "</td>"
            + "<td>" + dados[i].codbarras + "</td>"
            + "<td>" + dados[i].idmarca + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-primary' "
            + " onclick='alterar("+id+")'>Alterar</button>"
            + "<button type='button' class='btn btn-danger' "
            + " onclick='excluir("+id+")'>Excluir</button>"
            + "</td>"
            + "</tr>";
    }
}
function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/produto/" + idatual,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
                },
            method: "DELETE", 
            body: ""
        }
    ).then(() => {
        //fecha a dialog
        modalExcluir.hide();
        //recarrega a lista
        listar();
    })
}
function salvar() {
        const txtdescricao = document.getElementById("txtdescricao");
        const txtprecocusto = document.getElementById("txtprecocusto");
        const txtprecovenda = document.getElementById("txtprecovenda");
        const txtsaldoestoque = document.getElementById("txtsaldoestoque");
        const txtcodbarras = document.getElementById("txtcodbarras");
        const txtidmarca = document.getElementById("txtidmarca");

    const dados = {
        descricao: txtdescricao.value,
        precocusto: txtprecocusto.value,
        precovenda: txtprecovenda.value, 
        saldoestoque: txtsaldoestoque.value,
        codbarras: txtcodbarras.value, 
        idmarca: txtidmarca.value
    }

    var url;
    var metodo;
    if (idatual<=0) {
        //inserir
        url = "http://127.0.0.1:8080/produto";
        metodo = "POST";
    } else {
        //alterar
        url = "http://127.0.0.1:8080/produto/" + idatual;
        metodo = "PUT";
    }
    fetch(url,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
                },
            method: metodo, 
            body: JSON.stringify(dados)
        }
    ).then(() => {
        //abre a dialog
        modal.hide();
        //recarrega a lista
        listar();
    })

}

listar();
