var txtNome = document.getElementById("txtNome");
var txtCnpj = document.getElementById("txtCnpj");
var txtTelefone = document.getElementById("txtTelefone");
var txtCep = document.getElementById("txtCep");
var txtIdCidade = document.getElementById("idCidade");
var txtLogradouro = document.getElementById("txtLogradouro");
var txtNumero = document.getElementById("txtNumero");
var txtBairro = document.getElementById("txtBairro");

var txtPesquisa = document.getElementById("txtPesquisa");

const modal = new bootstrap.Modal(document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;

    //limpa os campo
    txtNome.value = "";
    txtCnpj.value = "";
    txtTelefone.value = "";
    txtCep.value = "";
    // txtIdCidade.value = "";
    txtLogradouro.value = "";
    txtNumero.value = "";
    txtBairro.value = "";

    //abre a dialog
    modal.show();

    document.getElementById("txtHeader").innerHTML = `<h1 class="modal-title fs-5" id="exampleModalLabel">Cadastrar fornecedor</h1>`
    document.getElementById("btnSalvar").innerHTML = `<button type="button" class="btn btn-primary" onclick="salvar()">Salvar</button>`;
};

function alterar(id) {
    idatual = id;
    //carregar os dados do id passado por parametro
    fetch("http://127.0.0.1:8080/fornecedor/" + id)
    .then(resp => resp.json())
    .then(dados => {
        //preenche os inputs

        txtNome.value = dados.nome;
        txtCnpj.value = dados.cnpj;
        txtTelefone.value = dados.telefone;
        txtCep.value = dados.cep;
        txtIdCidade.value = dados.idcidade;
        txtLogradouro.value = dados.logradouro;
        txtNumero.value = dados.numero;
        txtBairro.value = dados.bairro;

        //mostra a dialog para alterar
        modal.show();

        document.getElementById("txtHeader").innerHTML = `<h1 class="modal-title fs-5" id="exampleModalLabel">Alterar fornecedor</h1>`
        document.getElementById("btnSalvar").innerHTML = `<button type="button" class="btn btn-primary" onclick="salvarAlteracao(${idatual})">Salvar</button>`;
    });
};

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=11>Carregando...</td></tr>";

    // txtPesquisa.value != "" ? url = fetch("http://127.0.0.1:8080/fornecedor/" + txtPesquisa.value) : url = fetch("http://127.0.0.1:8080/fornecedor?pesquisa=")
    txtPesquisa.value != "" ? url = "http://127.0.0.1:8080/fornecedor/" + txtPesquisa.value : url = "http://127.0.0.1:8080/fornecedor";

    // fetch("http://127.0.0.1:8080/fornecedor?pesquisa=" + txtPesquisa.value)
    fetch(url)
    .then(resp => {
        // resp.json()
        if (resp.status == 200) {
            resp.json()
            .then(dados => {
                // console.log(dados)
                if (dados == "" || dados == null) {
                    const lista = document.getElementById("lista");
                    //limpa a lista
                    lista.innerHTML = "";
                    lista.innerHTML = "<tr><td colspan='9'>Nenhum registro encontrado</td></tr>";
                } else {
                    mostrar(dados, txtPesquisa.value);
                };
            });
        } else {
            const lista = document.getElementById("lista");
            //limpa a lista
            lista.innerHTML = "";
            lista.innerHTML = "<tr><td colspan='9'>Nenhum registro encontrado</td></tr>";
        };
    })
    .catch(err => {
        console.log("ERRO Catch: ", err)
    });
};

function mostrar(dados, pesquisa) {
    const lista = document.getElementById("lista");
    //limpa a lista
    lista.innerHTML = "";

    if (pesquisa == "") {
        for (var i in dados) {
            let id = dados[i].idfornecedor;
            lista.innerHTML += "<tr>"
                + "<td>" + id + "</td>"
                + "<td>" + dados[i].nome + "</td>"
                + "<td>" + dados[i].cnpj + "</td>"
                + "<td>" + dados[i].telefone + "</td>"
                + "<td>" + dados[i].cep + "</td>"
                + "<td>" + dados[i].idcidade + "</td>"
                + "<td>" + dados[i].logradouro + "</td>"
                + "<td>" + dados[i].numero + "</td>"
                + "<td>" + dados[i].bairro + "</td>"
                + "<td>"
                +   "<button type='button' class='btn btn-primary' onclick='alterar("+id+")'>Alterar</button>"
                + "</td>"
                + "<td>"
                +   "<button type='button' class='btn btn-danger' onclick='excluir("+id+")'>Excluir</button>"
                + "</td>"
                + "</tr>";
        }
    } else {
        let id = dados.idfornecedor;
        lista.innerHTML += "<tr>"
            + "<td>" + id + "</td>"
            + "<td>" + dados.nome + "</td>"
            + "<td>" + dados.cnpj + "</td>"
            + "<td>" + dados.telefone + "</td>"
            + "<td>" + dados.cep + "</td>"
            + "<td>" + dados.idcidade + "</td>"
            + "<td>" + dados.logradouro + "</td>"
            + "<td>" + dados.numero + "</td>"
            + "<td>" + dados.bairro + "</td>"
            + "<td>"
            +   "<button type='button' class='btn btn-primary' onclick='alterar("+id+")'>Alterar</button>"
            + "</td>"
            + "<td>"
            +   "<button type='button' class='btn btn-danger' onclick='excluir("+id+")'>Excluir</button>"
            + "</td>"
            + "</tr>";
    }
};

function excluir(id) {
    idatual = id;
    modalExcluir.show();
};

function excluirSim() {
    fetch("http://127.0.0.1:8080/fornecedor/" + idatual,
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
};

function salvar() {

    const dados = {
        nome: txtNome.value,
        cnpj: txtCnpj.value,
        telefone: txtTelefone.value,
        cep: txtCep.value,
        idcidade: Number(txtIdCidade.value),
        logradouro: txtLogradouro.value,
        numero: txtNumero.value,
        bairro: txtBairro.value
    };

    url = "http://127.0.0.1:8080/fornecedor";
    metodo = "POST";

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
};

function salvarAlteracao(id) {
    idatual = id;

    const dados = {
        nome: txtNome.value,
        cnpj: txtCnpj.value,
        telefone: txtTelefone.value,
        cep: txtCep.value,
        idcidade: Number(txtIdCidade.value),
        logradouro: txtLogradouro.value,
        numero: txtNumero.value,
        bairro: txtBairro.value
    };

    url = "http://127.0.0.1:8080/fornecedor/" + idatual;
    metodo = "PUT";

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
};

listar();