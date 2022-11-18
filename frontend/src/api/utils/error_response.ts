type ErrorResponse = {
    message: string,
    errors: Record<string, any>
}

export default ErrorResponse;