FROM mcr.microsoft.com/dotnet/core/sdk:2.2 AS build-env
WORKDIR /src/main/

COPY /src/main/*.csproj ./
RUN dotnet restore

COPY /src/main/. ./
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/core/aspnet:2.2
WORKDIR /src/main/
COPY --from=build-env /src/main/out .
EXPOSE 8080
ENTRYPOINT ["dotnet", "main.dll"]