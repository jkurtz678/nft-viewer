
const local_file_base_url = "http://localhost:8081/media/"

export const getLocalFileURL = (file_name: string): string => {
    return `${local_file_base_url}${file_name}`
}

// findLocalFile returns the filename of a local media file if it exists, otherwise returns an empty string
export const findLocalFile = async (token_id: string): Promise<string> => {
    const extensions = [".mp4", ".jpg"]

    for (const e of extensions) {
        try {
            const file_name = token_id + e;
            const res = await fetch(getLocalFileURL(file_name))
            if (res.status < 400) {
                return file_name;
            }
        } catch (err) {
            // do nothing
        }
    }

    return ""
}