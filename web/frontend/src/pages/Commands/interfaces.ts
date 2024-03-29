interface Command {
    name: string
    description: string
    response: string
    id?: string
    private?: boolean
    registered?: boolean
}
interface MapCommand {
    [key: string]: Command
}