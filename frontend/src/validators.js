function emailValidator () {
    return function email (value) {
        return (value && !!value.match(/[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/i)) || 'Please enter a valid email'
    }
}

function requiredValidator () {
    return function required (value) {
        return (value !== undefined && value !== null && value !== '') || 'This field is required'
    }
}

export {
    emailValidator,
    requiredValidator
}
