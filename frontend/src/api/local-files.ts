
const local_file_base_url = "http://localhost:8081/media/"

export const getLocalFileURL = (file_name: string) : string => {
    return `${local_file_base_url}${file_name}`
}

export const hasLocalFile = async (file_name: string) : Promise<boolean> => {
    try {
        const res = await fetch(getLocalFileURL(file_name))
        return res.status < 400;
    } catch (err) {
        return false
    }
}