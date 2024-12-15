import { defineEventHandler, readBody } from "h3";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const { weight, unit } = body;

  // Here, you would typically save the weight to your database
  // For example, using Prisma:
  // const savedWeight = await prisma.weight.create({
  //   data: { weight, unit, userId: event.context.auth.userId },
  // })

  // For now, we'll just return a mock response
  return {
    success: true,
    message: `Weight ${weight} ${unit} saved successfully`,
  };
});