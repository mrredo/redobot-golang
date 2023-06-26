interface Command {
    name: string
    description: string
    response: string
    id?: string
    registered?: boolean
}
interface MapCommand {
    [key: string]: Command
}