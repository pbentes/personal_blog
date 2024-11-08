import os
from urllib.parse import urljoin

base_url = "http://localhost:3000"  # Change to your site's base URL
output_dir = "docs"          # Change to your mdBook build directory

def generate_sitemap(output_dir, base_url):
    sitemap_entries = []
    for root, _, files in os.walk(output_dir):
        for file in files:
            if file.endswith(".html"):
                # Get relative path from the build directory and append to base URL
                rel_path = os.path.relpath(os.path.join(root, file), output_dir)
                url = urljoin(base_url, rel_path.replace(os.sep, '/'))
                sitemap_entries.append(f"<url><loc>{url}</loc></url>")

    sitemap_content = (
        "<?xml version='1.0' encoding='UTF-8'?>\n"
        "<urlset xmlns='http://www.sitemaps.org/schemas/sitemap/0.9'>\n"
        + "\n".join(sitemap_entries) +
        "\n</urlset>"
    )

    with open(os.path.join(output_dir, "sitemap.xml"), "w") as f:
        f.write(sitemap_content)

generate_sitemap(output_dir, base_url)