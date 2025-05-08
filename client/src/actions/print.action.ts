"use server"

export async function startAction(formData: FormData) {
    const name = formData.get("name")

    console.log(name)
    const data = JSON.stringify({
        name: name
    })

    const res = fetch("https://e742-46-7-24-170.ngrok-free.app/start", {
        method: "POST",
        body: data
    })
    
    console.log(res)
}

export async function stopAction(formData: FormData) {
    const name = formData.get("name")

    console.log(name)
    const data = JSON.stringify({
        name: name
    })

    const res = fetch("https://e742-46-7-24-170.ngrok-free.app/stop", {
        method: "POST",
        body: data
    })
    
    console.log(res)
}