import requests

count=0
def remove_empty_lines(text: str):
    lines = text.split('\n')
    non_empty_lines = [line for line in lines if line.strip() != '']
    cleaned_text = '\n'.join(non_empty_lines)
    return cleaned_text


# def split_code_into_chunks(code: str):
#     try:
#         code = remove_empty_lines(code)
#         lines = code.split('\n')
#         chunks = []
#         current_chunk = []
#         inside_class = False  # Flag to track if we are currently inside a class definition

#         for line in lines:
#             if line.startswith('class') or line.startswith('async class'):
#                 if current_chunk:
#                     chunks.append(current_chunk)
#                     current_chunk = []
#                 inside_class = True
#                 current_chunk.append(line)

#             elif line.startswith('def') or line.startswith('async def'):
#                 if current_chunk and not inside_class:
#                     chunks.append(current_chunk)
#                     current_chunk = []
#                 inside_class = False
#                 current_chunk.append(line)

#             elif line.startswith('#') or line.startswith('@'):
#                 continue

#             elif line.startswith("import ") or line.startswith("from "):
#                 continue
#             else:
#                 inside_class = False
#                 # inside_imports = False
#                 current_chunk.append(line)

#             # Check if the class definition has ended to reset the flag
#             if inside_class and line.strip().endswith(':'):
#                 inside_class = False

#         if current_chunk:
#             chunks.append(current_chunk)

#         non_empty_lists = [sublist for sublist in chunks if any(item.strip() != '' for item in sublist)]
#         return non_empty_lists
#     except Exception as e:
#         print(f"An error occurred in split_code_into_chunks: {e}")
#         return []


# def create_two_chunks(code_chunks):
#     result = []
#     current_chunk = []
#     inside = False
#     for line in code_chunks:
#         if line.startswith("def") or line.startswith("class") or line.startswith("async def ") or line.startswith(
#                 "async class "):
#             inside = True
#             if current_chunk:
#                 result.append(current_chunk)
#             current_chunk = [line]
#         elif line.startswith(' ') and inside == True:
#             current_chunk.append(line)
#         else:
#             if current_chunk:
#                 inside = False
#                 result.append(current_chunk)
#             current_chunk = []
#     if current_chunk:
#         result.append(current_chunk)
#     return result

def split_code_into_chunks(code):
    try:
        lines = code.split('\n')
        chunks = []
        current_chunk = []

        inside_function = False

        for line in lines:
            # Check for function declaration
            if line.strip().startswith('func '):
                if inside_function:
                    chunks.append(current_chunk)
                current_chunk = [line]
                inside_function = True
            elif inside_function:
                # Check for the end of a function
                if line.strip().endswith('}') and not line.strip().startswith('}'):
                    inside_function = False
                current_chunk.append(line)

        # Append the last chunk if not empty
        if current_chunk:
            chunks.append(current_chunk)

        # Remove empty and whitespace-only lines from chunks
        non_empty_lists = [sublist for sublist in chunks if any(item.strip() != '' for item in sublist)]
        return non_empty_lists

    except Exception as e:
        print(f"An error occurred in split_code_into_chunks: {e}")
        return []


# access_token = "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"


def get_branches(username, repository):
    api_url = f"https://api.github.com/repos/{username}/{repository}/branches"
    print(api_url)
    response = requests.get(api_url)

    if response.status_code == 200:
        branches = [branch['name'] for branch in response.json()]
        return branches
    else:
        print("Failed to fetch repository branches.")
        return []


def get_repository_files(username, repository, branch):
    api_url = f"https://api.github.com/repos/{username}/{repository}/contents?ref={branch}"
    print(api_url)
    response = requests.get(api_url)
    print(response)
    if response.status_code == 200:
        contents = response.json()
        all_code_content = []

        def fetch_content_recursive(items):
            for item in items:
                if item['type'] == 'file' and item['name'].endswith('.go'):
                    file_url = item['download_url']
                    print(file_url)
                    response = requests.get(file_url)
                    if response.status_code == 200:
                        content = response.text
                        all_code_content.append(content)
                    else:
                        print(f"Failed to fetch content for {item['name']}")
                elif item['type'] == 'dir':
                    subdir_url = item['url']
                    subdir_response = requests.get(subdir_url)
                    if subdir_response.status_code == 200:
                        subdir_contents = subdir_response.json()
                        fetch_content_recursive(subdir_contents)
                    else:
                        print(f"Failed to fetch contents for subdirectory {item['name']}")

        fetch_content_recursive(contents)
        return all_code_content
    else:
        print("Failed to fetch repository contents.")
        return []


def main(username, repository, selected_branch):
    global count
    branches = get_branches(username, repository)

    if selected_branch not in branches:
        print("Invalid branch name. Please select from the available branches:", branches)
        print(branches)

    code_contents = get_repository_files(username, repository, selected_branch)
    print(code_contents)
    combined_chunks = []
    code = []
    for code_content in code_contents:
        code.append(code_content.splitlines())
        chunks = split_code_into_chunks(code_content)
        print("chunks", chunks)
        print(len(chunks))
        count+=1
        print(count)
    #     for chunk in chunks:
    #         if any(line.startswith("class") or line.startswith("def") for line in chunk):
    #             codes = create_two_chunks(chunk)
    #             print("codes", codes)
    #             combined_chunks.extend(codes)
    #     combined_chunks.extend(code)
    # print("\nCombined Chunks:", combined_chunks)
    # print("\nCombined Chunks:")
    # for chunk in combined_chunks:
    #     if isinstance(chunk, list):
    #         print('chunk')
    #         print("\n".join(chunk))
    #     else:
    #         print(chunk)
    #     for idx, chunk in enumerate(chunks, start=1):
    #         codes.append(chunk)
    #         print(f"Chunk {idx}:\n")
    #         print('\n'.join(chunk))
    #         print("=" * 40)
    # print(len(combined_chunks))
    # return combined_chunks


if __name__ == "__main__":
    main("segmentio", "ksuid", "master")
