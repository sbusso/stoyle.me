<script>
    import { emailValidator, requiredValidator } from './validators.js'
    import { createFieldValidator } from './validation.js'
    import Clipboard from './Clipboard.svelte';

    const [ validity, validate ] = createFieldValidator(requiredValidator(), emailValidator())

    let url = "";
    let short = "";
    let newShort = "";
    let expires = 24;

    let loading = false;
    let availLoading = false;

    let shortIsAvailable = false;

    let csrf = document.getElementsByName("csrf_token")[0].content;

    const load = () => {
        let u = new URL(url)
        if (u.hostname.includes("stoyle.me")) {
            document.getElementById("domainErrorbtn").click();
            return
        }

        loading = true;
        fetch("/api/shorten", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': csrf,
            },
            credentials: "same-origin",
            body: JSON.stringify({url, short, expires})
        })
        .then((res) => (res.json()))
        .then((data) => {
            newShort = "https://stoyle.me/" + data.short;
            short = "";
            loading = false;
            shortIsAvailable = false;
        })
        .catch((err) => {
            console.error(err);
            loading = false;
        })
    };

    const checkAvailability = () => {
        availLoading = true;

        fetch("/api/checkavailability", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': csrf,
            },
            credentials: "same-origin",
            body: JSON.stringify({short})
        })
        .then((res) => (res.json()))
        .then((data) => {
            shortIsAvailable = data.message;
            availLoading = false;

            if (shortIsAvailable) {
                document.getElementById("short").style.borderBottomColor = "green";
                document.getElementById("short").style.borderBottomWidth = "medium";
            } else {
                document.getElementById("short").style.borderBottomColor = "red";
                document.getElementById("short").style.borderBottomWidth = "medium";
            }
        })
        .catch((err) => {
            console.error(err);
            availLoading = false;
        });
    };

    const copy = () => {
        const app = new Clipboard({
            target: document.getElementById('clipboard'),
            props: { newShort },
        });
        app.$destroy();
        document.getElementById("copyText").innerText = "copied";
    }
</script>

<style>
    #url {
        border: none;
        border-bottom: 2px solid #1361be;
    }
    #short {
        max-width: 20em;
    }
    input[type="text"]:focus,
    input[type="number"]:focus,
    button:focus {
        outline: none;
        box-shadow: none;
    }

    #short-btn {
        width: 100px;
    }
    #avail-btn {
        width: 150px;
    }
    .outer {
        background-color: #ffffff;
        padding-top: 2.5em;
        padding-bottom: 5em;
        box-shadow: #666666;
    }
    .card {
        text-align: center;
        padding-top: 5%;
    }
    .result {
        text-align: center;
        padding-top: 0;
        margin: 2%;
        max-width: 95%;
        background: transparent;
        font-size: 1.5em;
        border: none;
    }
    #expiry {
        width: 15%;
    }
    .navbar {
        margin-top: 5%;
    }
</style>

<nav class="navbar navbar-expand-lg navbar-dark container bg-dark">
    <a class="navbar-brand" href="/">stoyle.me</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="https://github.com/amalshaji/stoyle.me" target="_blank">GitHub</a>
            </li>
            <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Contact
                </a>
                <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                    <a class="dropdown-item" href="https://twitter.com/pydantic" target="_blank">Twitter @ pydantic</a>
                    <div class="dropdown-divider"></div>
                    <a class="dropdown-item" href="mailto:amalshajid@gmail.com">amalshajid@gmail.com</a>
                </div>
            </li>
            <li class="nav-item">
                <a class="nav-link" data-toggle="modal" data-target="#exampleModal" href="#">Why I created stoyle.me?
                </a>
            </li>
        </ul>
    </div>
</nav>

<!-- Modal for Why?  -->
<!-- Modal -->
<div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">Why I created stoyle.me?</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                Hello there👋
                <br><br>
                Hope you find this application useful. I'm a learning to build stuff and I built stoyle.me
                out of curiosity. I just wanted to know if I could do it.
                <br><br>
                The amount I spent on this(domain + hosting) is <strong>ZERO</strong>. The domain name was purchased with
                free credits from <strong>GitHub Student Developer Program</strong>. The hosting is done on GCP, because
                I had a lot of extra free credits lying around.
                <br><br>
                Thanks for taking your time to read this. I'm still learning, so you might find some bad coding practises.
                Please feel free to raise an <a href="https://github.com/amalshaji/stoyle.me/issues/new" target="_blank">issue</a> or a pull request.
                <br><br>
                And drop a star if you find this useful.
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-primary" data-dismiss="modal">Love your product</button>
                <button type="button" class="btn btn-success" data-dismiss="modal">Don't care</button>
            </div>
        </div>
    </div>
</div>

<button hidden id="domainErrorbtn" data-toggle="modal" data-target="#domainError"></button>

<div class="modal fade" id="domainError" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel1">Domain is locked</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                In order to prevent abuse, the <kbd>stoyle.me</kbd> domain is locked.
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-warning" data-dismiss="modal">chill dude</button>
            </div>
        </div>
    </div>
</div>


<div class="container outer shadow-lg p-3 mb-5 bg-white rounded">
    <form id="form" autocomplete="off">
        <!-- Define the URL input here -->
            <label for="url">Your URL</label>
            <div class="input-group mb-3">
                <input
                        type="text"
                        class="form-control"
                        id="url"
                        bind:value={url}
                        autofocus
                        use:validate={url}
                        aria-describedby="basic-addon3" />
                &nbsp;

            </div>

        <label for="basic-url">Expires in(hours)</label>
        <div id="expiry" class="form-group">
            <input type="number" bind:value={expires}  min="1" max="72" class="form-control" id="basic-url" >
        </div>
        <!-- Define the custom short input here -->
        <label for="short">Custom short</label>
        <div class="input-group mb-3">
            <div class="input-group-prepend">
                <span
                    class="input-group-text"
                    id="basic-addon3">https://stoyle.me/</span>
            </div>
            <input
                type="text"
                class="form-control"
                id="short"
                bind:value={short}
                aria-describedby="basic-addon3" />

        </div>
        <button
            type="button"
            class="btn btn-dark"
            id="avail-btn"
            on:click={checkAvailability}
            disabled={availLoading || short.length == 0}>
            {#if availLoading}
                <div
                    class="spinner-border spinner-border-sm text-light"
                    role="status"></div>
            {:else}Check Availability{/if}</button>
        <button
            type="button"
            class="btn btn-outline-primary waves-effect"
            on:click={load}
            disabled={loading || (!shortIsAvailable && short.length != 0) || !$validity.valid}
            id="short-btn">
            {#if loading}
                <div
                    class="spinner-border spinner-border-sm text-dark"
                    role="status"></div>
            {:else}Shorten{/if}
        </button>
    </form>


</div>

{#if newShort}
    <center>
        <div class="container result card">
            <div class="card-body">
                {newShort} <button class="btn btn-outline-info" id="copyText" on:click={copy}>copy</button>

            </div>
        </div>
    </center>
{/if}
<div id="clipboard"></div>
