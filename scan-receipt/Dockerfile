FROM mcr.microsoft.com/dotnet/core/sdk:2.2 AS build-env
WORKDIR /src/main/

# Copy csproj and restore as distinct layers
COPY /src/main/*.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY /src/main/. ./
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:2.2
WORKDIR /src/main/
COPY --from=build-env /src/main/out .
EXPOSE 8080
ENTRYPOINT ["dotnet", "main.dll"]