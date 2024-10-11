import { useRouteError } from 'react-router-dom';

export const UnexpectedError = () => {
    const error: unknown = useRouteError();

    return (
        <div>
            <h1>Unexpected Error had ocurred!</h1>
            <p>Sorry!</p>
            <p>{(error as Error)?.message ||
            (error as { statusText?: string })?.statusText}</p>
        </div>
    )
}