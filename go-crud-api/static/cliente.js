idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

function novo() {
    idatual = -1;
    const txtnome = document.getElementById("txtnome");
    const txtcpf = document.getElementById("txtcpf");
    const txtlogradouro = document.getElementById("txtlogradouro");
    const txtnumero = document.getElementById("txtnumero");
    const txtbairro = document.getElementById("txtbairro");
    const txtcep = document.getElementById("txtcep");
    const txttelefone = document.getElementById("txttelefone");
    const idcidade = document.getElementById("idcidade");

    //limpa os campo
    txtnome.value = "";
    txtcpf.value = "";
    txtlogradouro.value = "";
    txtnumero.value = "";
    txtbairro.value = "";
    txtcep.value = "";
    txttelefone.value = "";
    idcidade.value = "";
    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;

    fetch("http://127.0.0.1:8080/cliente/" + id)
        .then(resp => resp.json())
        .then(dados => {


            const txtnome = document.getElementById("txtnome");
            const txtcpf = document.getElementById("txtcpf");
            const txtlogradouro = document.getElementById("txtlogradouro");
            const txtnumero = document.getElementById("txtnumero");
            const txtbairro = document.getElementById("txtbairro");
            const txtcep = document.getElementById("txtcep");
            const txttelefone = document.getElementById("txttelefone");
            const idcidade = document.getElementById("idcidade");

            txtnome.value = dados.nome;
            txtcpf.value = dados.cpf;
            txtlogradouro.value = dados.logradouro;
            txtnumero.value = dados.numero;
            txtbairro.value = dados.bairro;
            txtcep.value = dados.cep;
            txttelefone.value = dados.telefone;
            idcidade.value = dados.idcidade;

            modal.show();
        });
}
function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    fetch("http://127.0.0.1:8080/cliente/" + txtpesquisa.value)
        .then(resp => resp.json())
        .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";

    for (var i in dados) {
        let id = dados[i].idcliente;
        lista.innerHTML += "<tr>"
            + "<td>" + id + "</td>"
            + "<td>" + dados[i].nome + "</td>"
            + "<td>" + dados[i].cpf + "</td>"
            + "<td>" + dados[i].logradouro + "</td>"
            + "<td>" + dados[i].numero + "</td>"
            + "<td>" + dados[i].bairro + "</td>"
            + "<td>" + dados[i].cep + "</td>"
            + "<td>" + dados[i].telefone + "</td>"
            + "<td>" + dados[i].idcidade + "</td>"

            + "<td style='display:flex'>"
            + "<button type='button' class='btn btn-primary' "
            + " onclick='alterar(" + id + ")' style='margin-right: 5px;'>Alterar</button>"
            + "<button type='button' class='btn btn-danger' "
            + " onclick='excluir(" + id + ")'>Excluir</button>"
            + "</td>"
            + "</tr>";
    }
}
function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/cliente/" + idatual,
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: "DELETE",
            body: ""
        }
    ).then(() => {

        modalExcluir.hide();
        listar();
    })
}
async function salvar() {
    const txtnome = document.getElementById("txtnome");
    const txtcpf = document.getElementById("txtcpf");
    const txtlogradouro = document.getElementById("txtlogradouro");
    const txtnumero = document.getElementById("txtnumero");
    const txtbairro = document.getElementById("txtbairro");
    const txtcep = document.getElementById("txtcep");
    const txttelefone = document.getElementById("txttelefone");
    const idcidade = document.getElementById("idcidade");
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");

    const dados = {
        nome: txtnome.value,
        cpf: txtcpf.value,
        logradouro: txtlogradouro.value,
        numero: txtnumero.value,
        bairro: txtbairro.value,
        cep: txtcep.value,
        telefone: txttelefone.value,
        idcidade: parseInt(idcidade.value)
    }

    var url;
    var metodo;
    if (idatual <= 0) {
        url = "http://127.0.0.1:8080/cliente";
        metodo = "POST";
    } else {
        url = "http://127.0.0.1:8080/cliente/" + idatual;
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
        modal.hide();
        listar();
    })

}

listar();