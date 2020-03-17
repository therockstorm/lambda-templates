import { APIGatewayProxyHandler } from "aws-lambda"
import "source-map-support/register"

export const node: APIGatewayProxyHandler = async (event, _context) => {
  return {
    statusCode: 200,
    body: JSON.stringify({ message: "response body", input: event }, null, 2)
  }
}
