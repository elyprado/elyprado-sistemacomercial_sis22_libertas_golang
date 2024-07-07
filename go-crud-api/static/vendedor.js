idatual = 0;

const modal = new bootstrap.Modal(
    document.getElementById('modalAlterar'));

const modalExcluir = new bootstrap.Modal(
    document.getElementById('modalExcluir'));

const modalTitle = document.getElementById("exampleModalLabel");

function novo() {
    idatual = -1;
    const txtnome = document.getElementById("txtnome");
    const txtcpf = document.getElementById("txtcpf");
    const txtlogradouro = document.getElementById("txtlogradouro");
    const txtnumero = document.getElementById("txtnumero");
    const txtbairro = document.getElementById("txtbairro");
    const txtcep = document.getElementById("txtcep");
    const txttelefone = document.getElementById("txttelefone");
    const perc_comissao = document.getElementById("perc_comissao");
    const idcidade = document.getElementById("idcidade");
    modalTitle.innerHTML = 'Cadastrar Vendedor <i class="fa-solid fa-user-pen"></i>'

    //limpa os campo
    txtnome.value = "";
    txtcpf.value = "";
    txtlogradouro.value = "";
    txtnumero.value = "";
    txtbairro.value = "";
    txtcep.value = "";
    txttelefone.value = "";
    perc_comissao.value = "";
    idcidade.value = "";

    //abre a dialog
    modal.show();
}
function alterar(id) {
    idatual = id;
    //carregar os dados do id passado por parametro
    fetch("http://127.0.0.1:8080/vendedor/" + id)
    .then(resp => resp.json())
    .then(dados => {
        //preenche os inputs
        const txtnome = document.getElementById("txtnome");
        const txtcpf = document.getElementById("txtcpf");
        const txtlogradouro = document.getElementById("txtlogradouro");
        const txtnumero = document.getElementById("txtnumero");
        const txtbairro = document.getElementById("txtbairro");
        const txtcep = document.getElementById("txtcep");
        const txttelefone = document.getElementById("txttelefone");
        const perc_comissao = document.getElementById("perc_comissao");
        const idcidade = document.getElementById("idcidade");
        modalTitle.innerHTML = 'Alterar Vendedor <i class="fa-solid fa-pen-to-square"></i>'

        txtnome.value = dados.nome;
        txtcpf.value = dados.cpf;
        txtlogradouro.value = dados.logradouro;
        txtnumero.value = dados.numero;
        txtbairro.value = dados.bairro;
        txtcep.value = dados.cep;
        txttelefone.value = dados.telefone;
        perc_comissao.value = dados.perc_comissao;
        idcidade.value = dados.idcidade;
        

        //mostra a dialog para alterar
        modal.show();
    });
}
function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan=5>Carregando...</td></tr>";

    const txtpesquisa = document.getElementById("txtpesquisa");


    
    fetch("http://127.0.0.1:8080/vendedor?pesquisa=" + txtpesquisa.value)
    .then(resp => resp.json())
    .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    //limpa a lista
    lista.innerHTML = "";
    //percorre os dados
    for (var i in dados) {
        let id = dados[i].idvendedor;
        lista.innerHTML += "<tr>"
            + "<td class='cID'>" + id + "</td>"
            + "<td class='cNOME'>" + dados[i].nome + "</td>"
            + "<td class='cCPF'>" + dados[i].cpf + "</td>"
            + "<td class='cLOG'>" + dados[i].logradouro + "</td>"
            + "<td class='cNUM'>" + dados[i].numero + "</td>"
            + "<td class='cBAIRRO'>" + dados[i].bairro + "</td>"
            + "<td class='cCEP'>" + dados[i].cep + "</td>"
            + "<td class='cTELEFONE'>" + dados[i].telefone + "</td>"
            + "<td class='cPERC'>" + dados[i].perc_comissao+ "%" +"</td>"
            + "<td class='cIDCID'>" + dados[i].idcidade + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-dark' "
            + " onclick='alterar("+id+")'><i class='fa-solid fa-pencil'></i></button></button>"
            + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-danger' style='margin-left:10px'"
            + " onclick='excluir("+id+")'><i class='fa-solid fa-trash'></i></button></button>"
            + "</td>"
            + "</tr>";
    }
}
function excluir(id) {
    idatual = id;
    modalExcluir.show();
}
function excluirSim() {
    fetch("http://127.0.0.1:8080/vendedor/" + idatual,
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
        const txtnome = document.getElementById("txtnome");
        const txtcpf = document.getElementById("txtcpf");
        const txtlogradouro = document.getElementById("txtlogradouro");
        const txtnumero = document.getElementById("txtnumero");
        const txtbairro = document.getElementById("txtbairro");
        const txtcep = document.getElementById("txtcep");
        const txttelefone = document.getElementById("txttelefone");
        const perc_comissao = document.getElementById("perc_comissao");
        const idcidade = document.getElementById("idcidade");

    const dados = {
        nome: txtnome.value,
        cpf: txtcpf.value,
        logradouro: txtlogradouro.value, 
        numero: txtnumero.value,
        bairro: txtbairro.value, 
        cep: txtcep.value,
        telefone: txttelefone.value, 
        perc_comissao: perc_comissao.value,
        idcidade: idcidade.value 
    }

    var url;
    var metodo;
    if (idatual<=0) {
        //inserir
        url = "http://127.0.0.1:8080/vendedor";
        metodo = "POST";
    } else {
        //alterar
        url = "http://127.0.0.1:8080/vendedor/" + idatual;
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

